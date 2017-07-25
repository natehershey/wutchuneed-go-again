import React, { Component } from 'react';

import Client from '../services/ApiClient'
import List from './List'
import ListsIndex from './ListsIndex'

import '../css/app.css';
import logo from '../images/cart.svg';

class App extends Component {
  state = {
    lists: [],
    currentList: {}
  };

  headerClass = () => {
    if (!(this.state.currentList && this.state.currentList.categories)) {
      return "show"
    }
    return "hide"
  }

  header = () => {
      return(
        <div className={"app-header " + this.headerClass()}>
          <img src={logo} className="app-logo" alt="logo" />
          <h1>Wutchuneed?</h1>
        </div>
      )
  }

  showResetButtonClass = () => {
    if (this.state.currentList && this.state.currentList.categories) {
      return "show"
    }
    return "hide"
  }

  resetCurrentList = () => {
    this.setState({
      currentList: null
    })
  }

  addCategory = (name, listId, e) => {
    // TODO: Validate name and listId
    Client.addCategory(name, listId, category => {
      Client.getList(this.state.currentList.id, list => {
        this.setState({
          currentList: list
        });
      });
    });
    this.resetInputField()
  }

  updateCategory = (categoryId, attributes) => {
    Client.updateCategory(categoryId, attributes, category => {
      Client.getList(this.state.currentList.id, list => {
        this.setState({
          currentList: list
        });
      });
    });
  }

  deleteCategory = (categoryId) => {
    Client.deleteCategory(categoryId, response => {
      if (response["success"] == true) {
        Client.getList(this.state.currentList.id, list => {
          this.setState({
            currentList: list
          });
        });
      } else {
        console.log("ERROR: unable to delete category ", categoryId)
      }
    });
  }

  resetInputField = (id) => {
    console.log("reset");
  }

  getList = (id) => {
    // TODO: Validate ID
    Client.getList(id, list => {
      this.setState({
        currentList: list
      });
    });
  }

  getLists = () => {
    Client.getLists(lists => {
      this.setState({
        lists: lists
      });
    });
  }

  componentDidMount() {
    console.log("componentDidMount()")
    this.getLists()
  }

  renderListsView() {
    return <ListsIndex lists={this.state.lists}
                        currentList={this.state.currentList}
                        onListClick={this.getList}
                        resetCurrentList={this.resetCurrentList}
                        showResetButtonClass={this.showResetButtonClass}/>
  }

  renderListDetailView() {
    // Don't bother rendering this if there's no list currently selected
    if (this.state.currentList && this.state.currentList.categories) {
      return <List list={this.state.currentList}
                    addCategoryHandler={this.addCategory}
                    deleteCategoryHandler={this.deleteCategory}
                    updateCategoryHandler={this.updateCategory}/>
    }
  }

  render() {
    const handleListClick = this.handleListClick
    return (
      <div className="app">
        {this.header()}
        <div className="list-chooser">
          {this.renderListsView()}
        </div>
        <div className="list-view">
          {this.renderListDetailView()}
        </div>
      </div>
    );
  }
}

export default App;
