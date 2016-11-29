package main

// Request .
type Request struct {
	Event     string      `json:"event"`
	Item      RequestItem `json:"item"`
	WebhookID int         `json:"webhook_id"`
}

// RequestItem .
type RequestItem struct {
	Message RequestMessage `json:"message"`
	Room    RequestRoom    `json:"room"`
}

// RequestMessage .
type RequestMessage struct {
	Date     string `json:"date"`
	From     User   `json:"from"`
	ID       string `json:"id"`
	Mentions []User `json:"mentions"`
	Message  string `json:"message"`
	Type     string `json:"type"`
}

// RequestRoom .
type RequestRoom struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// User .
type User struct {
	ID          int    `json:"id"`
	MentionName string `json:"mention_name"`
	Name        string `json:"name"`
}

// Example:
// {
//     event: 'room_message',
//     item: {
//         message: {
//             date: '2015-01-20T22:45:06.662545+00:00',
//             from: {
//                 id: 1661743,
//                 mention_name: 'Blinky',
//                 name: 'Blinky the Three Eyed Fish'
//             },
//             id: '00a3eb7f-fac5-496a-8d64-a9050c712ca1',
//             mentions: [],
//             message: '/weather',
//             type: 'message'
//         },
//         room: {
//             id: 1147567,
//             name: 'The Weather Channel'
//         }
//     },
//     webhook_id: 578829
// }
