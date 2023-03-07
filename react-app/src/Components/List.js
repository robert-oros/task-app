import Card from './Card';
import '../css/List.css'
import React from 'react';

let dragged
class List extends React.Component {
  constructor(){
    super()
    // this.state = {
    //   dragged: false
    // }
    this.dragStart = this.dragStart.bind(this)
  }

  dragStart(e) {
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
    console.log(e.target)

    if (e.target.classList.contains("dropzone")) {
      let parentOfTargetElem = e.target.parentNode.parentNode

      let elemToRm = parentOfTargetElem.parentNode.parentNode.parentNode.querySelector(".dragging")
      let parentOfElemToRm = elemToRm.parentNode
      parentOfElemToRm.removeChild(elemToRm)
      console.log(dragged, parentOfElemToRm)

      e.target.classList.remove("dragover");

      parentOfTargetElem.appendChild(dragged);
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
            <div className="card-list dropzone" onDragOver={(e) => this.dragOver(e)}> 
              {cards}
            </div>
            <button className='todo-add' type="button">Adauga Todo</button>
          </div>
        </div>
    );
  }
}
  
export default List;