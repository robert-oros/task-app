import '../css/Card.css'
import React, { Component } from 'react';

class Card extends Component {
  render() {
    return (
      <ul 
        draggable
        className="card-text-container dropzone"
        onDragStart={(e) => this.props.dragStart(e)}
        onDragEnd={(e) => this.props.dragEnd(e)}
        onDragOver={(e) => this.props.dragOver(e)}
        onDragEnter={(e) => this.props.dragEnter(e)}
        onDragLeave={(e) => this.props.dragLeave(e)}
        onDrop={(e) => this.props.drop(e)}
      >
        <li className='card-id'>
          <span className="label">{this.props.data.cardId}</span>
        </li>
        <li className='dropzone'>
          <span className="card-text">{this.props.data.text}</span>
        </li>
      </ul>
    );
  }
}

export default Card;