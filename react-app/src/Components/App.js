import React, { Component } from 'react';
import Board from './Board';
import Navbar from './Navbar'

class App extends Component {
  constructor(){
    super()
    this.state = {
      boards: [],
      isLoaded: false
    }
  }

  componentDidMount(){
    fetch("http://localhost:8081/get_boards", {
      method: 'GET'
    })
    .then(res => res.json())
    .then(data => {
      this.setState({ boards: data})
      this.setState({ isLoaded: true })
    })
  }
  
  render() {
    let boards;
    if (typeof this.state.boards !== 'undefined') {
      boards = this.state.boards?.map(b => {
        return <Board data={b} />
      })
    }

    if (this.state.isLoaded) {
      return (
        <div>
          <Navbar />
          {boards}
        </div>
      );
    } else {
      return (
        <h4 style={{textAlign: "center", paddingTop: '50%', paddingBottom: '50%'}}>Loading Board...</h4>
      );
    }
  }
}

export default App;
