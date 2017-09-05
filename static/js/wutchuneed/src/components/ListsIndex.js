import React from "react";
import PropTypes from 'prop-types';

import upToListIcon from '../images/chevron-circle-up.svg';

import save from '../images/check-circle.svg';
import cancel from '../images/x-mark-circle.png';

import '../css/lists_index.css';

class ListsIndex extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      addingList: true,
      newListName: ''
    };
  }

  updateNewListName = (e) => {
    this.setState({
      newListName: e.target.value
    });
  }

  selectedClass(id) {
    // console.log("id: ", id)
    // console.log("currentListId: ", this.props.currentList.id)
    if (!(this.props.currentList && this.props.currentList.categories)) {
      return "inactive"
    } else {
      return (id === this.props.currentList.id) ? 'active' : 'inactive hide'
    }
  }

  showAddListFormClass() {
    if (this.props.currentList && this.props.currentList.categories) {
      return "hide"
    }
  }

  listViewActiveClass() {
    if (this.props.currentList && this.props.currentList.categories) {
      return "list-view-active"
    }
    return "list-view-inactive"
  }

  clearListInputs() {

  }

  handleNewListSubmit = () => {
    const newListName = this.state.newListName;
    const newListType = this.state.newListType;

    this.props.addListHandler(newListName, newListType);
    // this.clearListInputs();
  }

  render() {
    const listRows = this.props.lists.map((list, idx) => (
      <div key={idx}
        className={"list-index-row " + this.selectedClass(list.id)}
        onClick={() => { this.props.onListClick(list.id) }}>
        {list.name}
      </div>
    ));

    return (
      <div className={"lists-container " + this.listViewActiveClass()}>
        <div className={"add-list " + this.showAddListFormClass()}>
          <form onSubmit={this.handleNewListSubmit}>
            <div className="new-list">
              <div className="new-list-input-container">
                <input type="text"
                        id="new-list-input-field"
                        value={this.state.newListName}
                        onChange={(e) => this.updateNewListName(e)} />
              </div>
              <div className="new-list-save-button">
                <img className="new-list-save-image"
                      alt="Save new list"
                      src={save}
                      onClick={this.handleNewListSubmit}/>
              </div>
              <div className="new-list-cancel-button">
                <img className="new-list-cancel-image"
                      alt="Cancel new list"
                      src={cancel}
                      onClick={this.clearListInputs}/>
              </div>
            </div>
          </form>
        </div>
        <div className="list-index">
          {listRows}
        </div>
        <div className={"list-index-reset-button " + this.props.showResetButtonClass()}
              onClick={this.props.resetCurrentList}>
          <img src={upToListIcon}
                className="reset-button-image"
                alt="show all lists" />
        </div>
      </div>
    );
  }
}

ListsIndex.propTypes = {
  addListHandler: PropTypes.func.isRequired,
  lists: PropTypes.array.isRequired,
  currentList: PropTypes.object.isRequired,
  onListClick: PropTypes.func.isRequired,
  resetCurrentList: PropTypes.func.isRequired,
  showResetButtonClass: PropTypes.func.isRequired
};


export default ListsIndex;
