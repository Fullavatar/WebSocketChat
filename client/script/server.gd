extends Control


func _on_connect_pressed() -> void:
	get_tree().call_group("login", "serverConnect", $Address.text, $Port.text)
