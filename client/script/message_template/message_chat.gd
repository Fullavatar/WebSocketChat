extends Message
class_name Chat

@export var content:String


func _init() -> void:
	type = Global.MessageType.CHAT


func to_json() -> String:
	var message_dict = {
		"Type": type_to_string(),
		"Content": content,
	}
	return JSON.stringify(message_dict, "\t")
