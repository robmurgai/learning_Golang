// Unfortunately, it’s not always possible to send a message immediately because the relay satellite is not always above the horizon. Implement a
// buffer goroutine that receives messages sent from the rover and buffers them into a slice until they can be sent back to Earth.
//
// Implement Earth as a goroutine that receives messages only occasionally (in reality for a couple of hours every day, but you might want to make
// the interval a little shorter than that). Each message should contain the coordinates of the cell where the life might have been found, and the
// life value itself.

package practicefundamentals

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type emt struct {
	messageNum int
	message    string
	rsrev      string
	rssend     string
	esr        string
}

var msgTstr []emt

type earthMsg struct {
	messageNum int
	message    string
}

type earthMsgPkt struct {
	messages []earthMsg
	mu       sync.Mutex
}

// NewEarthMessager initializes Earth satelites and buffers for messages.
func NewEarthMessager() *chan string {
	var msgPacket earthMsgPkt
	upstreamMsgChan := make(chan string)
	earthMsgChan := make(chan earthMsg)

	go relaySatelliteReceiving(&msgPacket, upstreamMsgChan)
	go relaySatelliteSending(&msgPacket, earthMsgChan)
	go earthSatellite(earthMsgChan)
	//log.Printf("DEBUG: NewEarthMessager(): Earth Satelites Initialized\n")

	return &upstreamMsgChan

}

func relaySatelliteReceiving(msgPacket *earthMsgPkt, upstreamMsgChan chan string) {

	//log.Printf("DEBUG: relaySatelliteReceiving(): Initialized\n")
	var msgNumReceived = 0

	for msg := range upstreamMsgChan {
		//Lock the msgPackage while we are writing to it.
		msgPacket.mu.Lock()
		message := earthMsg{messageNum: msgNumReceived, message: msg}

		defer func() {
			msgPacket.mu.Unlock()
			if r := recover(); r != nil {
				log.Printf("ERROR: relaySatelliteReceiving(): When receiving Msg[%v]: %v", message.messageNum, message.message)
				log.Printf("ERROR: relaySatelliteReceiving(): msgNumReceived: %v", msgNumReceived)
				log.Printf("ERROR: relaySatelliteReceiving(): len(msgPacket.messages):%v and cap(msgPacket.messages):%v\n", len(msgPacket.messages), cap(msgPacket.messages))
				for _, value := range msgTstr {
					fmt.Printf("%v\n", value)
				}
				for _, value := range msgPacket.messages {
					log.Printf("%45v [%v]%v\n", " ", value.messageNum, value.message)
				}
				log.Printf("ERROR: relaySatelliteReceiving(): %+v\n", r)
				os.Exit(1)
			}
		}()

		msgPacket.messages = append(msgPacket.messages, message)
		msgNumReceived++
		msgTstr = append(msgTstr, emt{message.messageNum, message.message[:17], "Y", "", ""})
		log.Printf("INFO: relaySatelliteReceiving():  Msg[%v]: %v", message.messageNum, message.message)
		//log.Printf("DEBUG: Step 1 of 3: %v", msgTstr[message.messageNum])
		msgPacket.mu.Unlock()
	}

}

func relaySatelliteSending(msgPacket *earthMsgPkt, earthMsgChan chan earthMsg) {
	//log.Printf("DEBUG: relaySatelliteSending(): Initialized\n")
	for {
		msgPacket.mu.Lock()
		defer msgPacket.mu.Unlock()

		tmpMessages := msgPacket.messages
		for msgNumSent, message := range tmpMessages {

			defer func() {
				if r := recover(); r != nil {
					log.Printf("ERROR: relaySatelliteSending(): When sending Msg[%v]: %v", message.messageNum, message.message)
					for _, value := range msgPacket.messages {
						log.Printf("%v\n", value)
					}
					for _, value := range msgTstr {
						log.Printf("%v\n", value)
					}
					log.Printf("ERROR: relaySatelliteSending(): %+v\n", r)
					os.Exit(1)
				}
			}()

			//Confirm we are not resending messages
			if msgTstr[message.messageNum].rssend == "Y" {
				panic(fmt.Sprintf("Trying to resend msg: %v", message.messageNum))
			}

			earthMsgChan <- message
			msgTstr[message.messageNum].rssend = "Y"
			log.Printf("INFO: relaySatelliteSending():    Msg[%v]: %v", message.messageNum, message.message)
			//log.Printf("DEBUG: Step 2 of 3: %v", msgTstr[message.messageNum])

			// Remove the sent message from the global message queue. Lock the msgPackage while we are writing to it.
			msgPacket.messages = tmpMessages[msgNumSent+1:]

			// Was message.messageNum removed
			// check if any of the meesages in msgPackat have the message Number message.messageNum
			for _, tmpMsg := range msgPacket.messages {
				if tmpMsg.messageNum == message.messageNum {
					log.Printf("ERROR: relaySatelliteSending(): %v\n", tmpMessages[0])
					panic(fmt.Sprintf("Sent msg: %v not removed from msgPacket", tmpMsg))
				}
			}
		}
		msgPacket.mu.Unlock()
	}
}

// Implement Earth as a goroutine that receives messages only occasionally (in reality for a couple of hours every day, but you might want to make
// the interval a little shorter than that). Each message should contain the coordinates of the cell where the life might have been found, and the
// life value itself.
// Can receive messages for 1 second, every 5 seconds
func earthSatellite(upstreamMsgChan chan earthMsg) {
	//log.Printf("DEBUG: earthSatellite(): Initialized\n")
	msgRecievingWindowClosed := time.After(time.Second) // This is the window that could be 2 hours every 24 hours.  For now its 1 sec long
	for {
		//log.Printf("DEBUG: earthSatellite(): Message Window Availabale")
		select {
		case <-msgRecievingWindowClosed:
			//log.Printf("DEBUG: earthSatellite(): Message Window Unavailable")
			time.Sleep(5 * time.Second)                        // This is the time earth cannot receive messages, it might be 22 hours in real life.
			msgRecievingWindowClosed = time.After(time.Second) // This is the window that could be 2 hours every 24 hours  For now its 1 sec
		case message := <-upstreamMsgChan:
			if msgTstr[message.messageNum].esr == "Y" {
				log.Printf("ERROR: EarthSatellite(): msgTstr[%v] has already been received.", message.messageNum)
				log.Printf("ERROR: EarthSatellite(): msgTstr")
				for _, value := range msgTstr {
					fmt.Printf("%v\n", value)
				}
				panic(fmt.Sprintf("ERROR: EarthSatellite(): Trying to resend msg: %v", message.messageNum))
			}
			msgTstr[message.messageNum].esr = "Y"
			log.Printf("INFO: earthSatellite(): Received: Msg[%v]: %v", message.messageNum, message.message)
			//log.Printf("DEBUG: Step 3 of 3: %v", msgTstr[message.messageNum])
		}
	}
}

func (msgTstr emt) String() string {

	// messageNum int
	// message    string
	// rsrev      string
	// rssend     string
	// esr        string

	return fmt.Sprintf("(%-2v) %16v %2v %2v %2v", msgTstr.messageNum, msgTstr.message, msgTstr.rsrev, msgTstr.rssend, msgTstr.esr)
}
