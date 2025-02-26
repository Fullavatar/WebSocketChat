extends Message
class_name Connection

@export var username:String
@export var password:String


func _init() -> void:
	type = Global.MessageType.CONNECTION


func to_json() -> String:
	var message_dict = {
		"Type": type_to_string(),
		"Username": username,
		"Password": password,
	}
	return JSON.stringify(message_dict, "\t")
