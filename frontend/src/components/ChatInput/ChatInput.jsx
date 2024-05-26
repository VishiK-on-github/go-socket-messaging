import React, {Component} from "react";
import "./ChatInput.scss"

class ChatInput extends Component {
  render() {
    return (
      <div className="chat-input pb-4">
        <input onKeyDown={this.props.send} placeholder="Send Message" />
      </div>
    )
  }
}

export default ChatInput