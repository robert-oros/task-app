// import { Component } from "react";
import React, {Component} from "react";

class AddInput extends Component {
    constructor(){
        super()
        this.state = {
            name: ""
        }
        this.handleChange = this.handleChange.bind(this)
    }
    handleChange(event) {
        const {name, value} = event.target
        this.setState({
            name: value
        })
    }

    addBoard(name){
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

    render(){
        return (
            <div>
                <input type="text" value={this.state.name} placeholder="Border Name" onChange={this.handleChange}/>
                <input onClick={this.addBoard(this.state.name)} type="submit"/>
            </div>
        )
    }
}

export default AddInput