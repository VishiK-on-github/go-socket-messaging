import React, {Component} from "react";
import Message from "../Message";

class ChatHistory extends Component {
  render() {
    const messages = this.props.chatHistory.map(msg => <Message key={msg.timeStamp} message={msg.data} />)

    return (
      <div className="p-3">
        <h2 className="p-3">Chat History</h2>
        <div className="p-3">
          {messages}
        </div>
      </div>
    )
  }
}

export default ChatHistory