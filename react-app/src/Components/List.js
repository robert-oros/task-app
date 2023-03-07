import Card from './Card';
import '../css/List.css'
import React from 'react';

let dragged

class List extends React.Component {
  constructor(){
    super()
  }

  dragStart = (e) => {
    dragged = e.target
    e.target.classList.add("dragging");
  }
    
  dragEnd = (e) => {
    e.target.classList.remove("dragging");
    this.setState({
      dragged: e.target
    })
  }
    
  dragOver = (e) => {
    e.preventDefault();
  }
  
  dragEnter = (e) => {
    if (e.target.classList.contains("dropzone")) {
      // afterBeforElem = e.target
      e.target.classList.add("dragover");
    }
  }

  dragLeave = (e) => {
    if (e.target.classList.contains("dropzone")) {
      e.target.classList.remove("dragover");
    }
  }
  
  drop = (e) => {
    e.preventDefault();
    let parentOfTarget, parentOfDragged, elemToRemove;

    if (e.target.hasChildNodes()) {
      parentOfTarget = e.target.parentNode.parentNode
      elemToRemove = parentOfTarget.parentNode.parentNode.parentNode.querySelector(".dragging")
      
      parentOfDragged = elemToRemove.parentNode
      parentOfDragged.removeChild(elemToRemove)
      
      e.target.classList.remove("dragover");
      parentOfTarget.appendChild(dragged);
    } else {
      if (e.target.classList.contains("dropzone")) {
        parentOfTarget = e.target
        elemToRemove = parentOfTarget.parentNode.parentNode.parentNode.querySelector(".dragging")
        
        parentOfDragged = elemToRemove.parentNode
        parentOfDragged.removeChild(elemToRemove)
        e.target.classList.remove("dragover");
      }
    }
  }

  render() {
    const cards = this.props.data.cards.map(c => {
      return <Card 
        data={c} 
        dragStart={this.dragStart} 
        dragEnd={this.dragEnd}
        dragOver={this.dragOver}
        dragEnter={this.dragEnter}
        dragLeave={this.dragLeave}
        drop={this.drop}
      />
    })

    // addTodo() {

    // }

    // deleteTodo() {

    // }
    return (
      <div class="col list-container">
        <ul className='list-text'>
          {/* <li>
            <span className="label">{props.data.listId}</span>
          </li> */}
          <li>
            <span className="label">{this.props.data.title}</span>
          </li>
        </ul>

        <div>
          <div className="card-list dropzone" 
            onDrop={(e) => this.drop(e)}
            onDragOver={(e) => this.dragOver(e)}> 
            {cards}
          </div>
          <button className='todo-add' type="button">Adauga Todo</button>
        </div>
      </div>
    );
  }
}
  
export default List;