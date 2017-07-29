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
    return "hidden"
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
      if (response["success"] === true) {
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

  addItem = (attributes) => {
    console.log("Adding item: ", attributes)
    // TODO: Validate input
    Client.addItem(attributes, category => {
      Client.getList(this.state.currentList.id, list => {
        this.setState({
          currentList: list
        });
      });
    });
  }

  updateItem = (itemId, attributes) => {
    Client.updateItem(itemId, attributes, item => {
      Client.getList(this.state.currentList.id, list => {
        this.setState({
          currentList: list
        });
      });
    });
  }

  deleteItem = (itemId) => {
    Client.deleteItem(itemId, response => {
      if (response["success"] === true) {
        Client.getList(this.state.currentList.id, list => {
          this.setState({
            currentList: list
          });
        });
      } else {
        console.log("ERROR: unable to delete item ", itemId)
      }
    });
  }

  resetInputField = (id) => {
    console.log("reset");
  }


  componentDidMount() {
    console.log("componentDidMount()")
    this.getLists()
  }

  header = () => {
      return(
        <div className={"app-header " + this.headerClass()}>
          <img src={logo} className="app-logo" alt="logo" />
          <div className="app-name">Wutchuneed?</div>
        </div>
      )
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
                    updateCategoryHandler={this.updateCategory}
                    addItemHandler={this.addItem}
                    updateItemHandler={this.updateItem}
                    deleteItemHandler={this.deleteItem}/>
    }
  }

  render() {
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
