// import { Component } from "react";
import React, {Component} from "react";

function addBoard(name){
    fetch("http://localhost:8081/add_board", {
      method: 'POST',
      body: JSON.stringify({
        Name: name,
      }),
    })
    .then(res => console.log(res))
    .catch(error => {
      console.error(error)
    })
}

class AddInput extends Component {
    constructor(){
        super()
        this.state = {
            value: "",
            isLoadding: false
        }
        this.handleChange = this.handleChange.bind(this)
        this.setStateforLoadding = this.setStateforLoadding.bind(this)
    }
    handleChange(event) {
        this.setState({
            value: event.target.value
        })
    }

    setStateforLoadding(){
        this.setState({
            isLoadding: true
        })
    }

    render(){
        return (
            <div>
                <input type="text" value={this.state.value} placeholder="Border Name" onChange={this.handleChange}/>
                <input type="submit" value="Submit" onClick={this.setStateforLoadding}/>
                {this.state.isLoadding ? addBoard(this.state.value) : <div></div>}
          
            </div>
        )
    }
}

export default AddInput