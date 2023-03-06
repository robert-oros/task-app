import Card from './Card';
import '../css/List.css'
import React from 'react';

class List extends React.Component {
  constructor(){
    super()
    this.state = {
      dragged: false
    }
    this.dragStart = this.dragStart.bind(this)
  }

  dragStart(e) {
    this.setState({
      dragged: e.target
    })

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
    console.log("enter", e.target)
    // if (e.target.classList.contains("dropzone")) {
    //   e.target.classList.add("dragover");
    // }
  }

  dragLeave = (e) => {
    console.log("leave", e.target)
    // if (e.target.classList.contains("dropzone")) {
    //   e.target.classList.remove("dragover");
    // }
  }
  
  drop = (e) => {
    e.preventDefault();

    if (e.target.classList.contains("dropzone-li")) {
      let parent = e.target.parentNode.parentNode
      console.log(this.state.dragged)
      // let elemToRm = parent.querySelector(".dragging")
      
      // e.target.classList.remove("dragover");
      // console.log(parent)
      // document.body.remove(elemToRm)
      // e.target.parentNode.parentNode.parentNode.parentNode.appendChild(this.state.dragged);
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

          <div className="card-list">
            {cards}

            <ul className="card-text-container">
              <button type="button">Adauga Todo</button>
            </ul>

          </div>
        </div>
    );
  }
}
  
export default List;