extends Resource
class_name Message

@export var type: Global.MessageType


func to_json() -> String:
	var message_dict = {
		"Type": type_to_string(),
	}
	return JSON.stringify(message_dict, "\t")


func type_to_string() -> String:
	match type:
		Global.MessageType.SIGNUP: return "SignUp"
		Global.MessageType.FORGET: return "Forget"
		Global.MessageType.CONNECTION: return "Connection"
		Global.MessageType.DISCONNECTION: return "Disconnection"
		Global.MessageType.PING: return "Ping"
		Global.MessageType.CHAT: return "Chat"
		Global.MessageType.PRIVATE_CHAT: return "PrivateChat"
		Global.MessageType.STATUS: return "Status"
		Global.MessageType.ERROR: return "Error"
		Global.MessageType.ROOM_JOIN: return "RoomJoin"
		Global.MessageType.ROOM_LEAVE: return "RoomLeave"
	return "Unknown"
