import React from "react";
import PropTypes from 'prop-types';

import upToListIcon from '../images/chevron-circle-up.svg';

import '../css/lists_index.css';

class ListsIndex extends React.Component {
  selectedClass(id) {
    if (!(this.props.currentList && this.props.currentList.categories)) {
      return "inactive"
    } else {
      return (id === this.props.currentList.id) ? 'active' : 'inactive hide'
    }
  }

  listViewActiveClass() {
    if (this.props.currentList && this.props.currentList.categories) {
      return "list-view-active"
    }
    return "list-view-inactive"
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
        <div className="list-index">
          {listRows}
        </div>
        <div className={"list-index-reset-button " + this.props.showResetButtonClass()} onClick={() => this.props.resetCurrentList()}>
          <img src={upToListIcon} className="reset-button-image" alt="show all lists" />
        </div>
      </div>
    );
  }
}

ListsIndex.propTypes = {
  lists: PropTypes.array.isRequired,
  currentList: PropTypes.object,
  onListClick: PropTypes.func.isRequired,
  resetCurrentList: PropTypes.func.isRequired,
  showResetButtonClass: PropTypes.func.isRequired
};


export default ListsIndex;
