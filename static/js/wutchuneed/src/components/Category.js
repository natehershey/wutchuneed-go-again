import React from "react";
import PropTypes from 'prop-types';
import ReactConfirmAlert, { confirmAlert } from 'react-confirm-alert';

import Item from '../components/Item'
import ItemForm from '../components/ItemForm'

import moreImage from '../images/more-circle.png';
import editImage from '../images/edit-circle.svg';
import deleteImage from '../images/x-mark-circle.png';
import checkmarkImage from '../images/check-circle.svg';
import plusImage from '../images/character-sign-add-increase-plus-math-icon.svg';

import 'react-confirm-alert/src/react-confirm-alert.css'
import '../css/category.css';

class Category extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      forceHideChildren: false,
      autoHideChildren: false,
      inEditMode: false,
      editingCategoryName: false,
      newCategoryName: "",
      showDialog: false,
      addingItem: false,
      gotEmAll: false,
      hideOverride: false
    }
  }

  toggleForceHideChildren = () => {
    this.setState({
      forceHideChildren: !this.state.forceHideChildren,
    })
  }

  toggleEditingCategory = () => {
    this.setState({
      inEditMode: !this.state.inEditMode
    });
  }

  toggleEditingCategoryName = () => {
    this.setState({
      editingCategoryName: !this.state.editingCategoryName
    });
  }

  updateNewCategory = (e) => {
    this.setState({
      newCategoryName: e.target.value
    });
  }

  categoryNameFieldstrikethroughClass = () => {
    if (this.props.category.items.every(x => x.status === "in_cart") && this.props.category.items.length > 0) {
      return "strikethrough"
    }
  }

  categoryEditFieldsShowHideClass = () => {
    if (this.state.inEditMode === true && this.state.editingCategoryName === false) {
      return "show"
    }
    return "hide"
  }

  addItemButtonShowHideClass = () => {
    if (this.state.inEditMode === true) {
      return "hide"
    }
    return "show"
  }

  editCategoryButtonShowHideClass = () => {
    if (this.state.editingCategoryName === false) {
      return "show"
    }
    return "hide"
  }

  categoryNameShowHideClass = () => {
    if (this.state.editingCategoryName === true) {
      return "hide"
    }
    return "show"
  }

  editCategoryInputShowHideClass = () => {
    if (this.state.editingCategoryName === true) {
      return "show"
    }
    return "hide"
  }

  deleteCategoryButtonShowHideClass = () => {
    if (this.state.inEditMode === true && this.state.editingCategoryName === false) {
      return "show"
    }
    return "hide"
  }

  editCategoryNameButtonShowHideClass = () => {
    if (this.state.inEditMode === true && this.state.editingCategoryName === false) {
      return "show"
    }
    return "hide"
  }

  saveNewCategoryNameButtonShowHideClass = () => {
    if (this.state.editingCategoryName === true) {
      return "show"
    }
    return "hide"
  }

  cancelNewCategoryNameButtonShowHideClass = () => {
    if (this.state.editingCategoryName === true) {
      return "show"
    }
    return "hide"
  }

  deleteCategoryAlert = (id) => {
    const titleText = "Are you sure?"
    const messageText = "If you delete this category, the items go too."
    confirmAlert({
      title: titleText,                        // Title dialog
      message: messageText,               // Message dialog
      childrenElement: () => null,       // Custom UI or Component
      confirmLabel: 'Do it!',                           // Text button confirm
      cancelLabel: 'Cancel',                             // Text button cancel
      onConfirm: () => this.deleteCategory(id),    // Action after Confirm
      onCancel: () => {},      // Action after Cancel
    })
  }

  deleteCategory = (id) => {
    this.props.deleteCategoryHandler(id)
  }

  // TODO: 500 error when submitting a new name twice
  updateCategoryName = (categoryId) => {
    this.props.updateCategoryHandler(categoryId, {name: this.state.newCategoryName}, category => {
      // console.log("callback")
    });
    this.setState({
      newCategoryName: "",
      inEditMode: false,
      editingCategoryName: false
    });
  }

  cancelCategoryNameUpdate = () => {
    this.setState({
      editingCategoryName: false
    });
  }

  addItem = () => {
    this.setState({
      addingItem: true
    });
  }

  closeAddItemDialog = () => {
    this.setState({
      addingItem: false
    });
  }

  items() {
    if(!(this.state.forceHideChildren || this.state.autoHideChildren)) {
      return this.props.category.items.map((item, idx) => (
        <Item key={"item-" + item.id}
              item={item}
              deleteItemHandler={this.props.deleteItemHandler}
              updateItemHandler={this.props.updateItemHandler}/>
      ));
    }
  }

  itemForm() {
    if (this.state.addingItem === true) {
      return (
        <div className="add-item-dialog">
          <ItemForm list={this.props.list}
                    category={this.props.category}
                    handleClose={this.closeAddItemDialog}
                    addItemHandler={this.props.addItemHandler}
                    closeMe={this.closeAddItemDialog} />
        </div>
      )
    }
  }

  render() {
    const { category } = this.props;

    return (
      <div key={"cat-" + category.id} className={"category-container"}>
        <div className="category-row">

          <div className={"category-name " + this.categoryNameShowHideClass() + " " + this.categoryNameFieldstrikethroughClass()}
                onClick={() => this.toggleForceHideChildren()}>
            {category.name}
          </div>

          <div className={"edit-category-name " + this.editCategoryInputShowHideClass()}>
            <input type="text"
                    className="edit-category-name-input"
                    value={this.state.newCategoryName}
                    onChange={this.updateNewCategory} />
          </div>

          <div className={"add-item-to-category-button " + this.addItemButtonShowHideClass()}>
            <img className={"add-item-to-category-image button-image"}
                  alt="Add an item to this category"
                  src={plusImage}
                  onClick={this.addItem} />
          </div>

          <div className={"delete-category-button " + this.deleteCategoryButtonShowHideClass()}>
            <img className="delete-category-image button-image"
                  alt="Delete this category"
                  src={deleteImage}
                  onClick={() => this.deleteCategoryAlert(category.id)} />
          </div>

          <div className={"edit-category-name-button " + this.editCategoryNameButtonShowHideClass()}>
            <img className="edit-category-name-image button-image"
                  alt="Ecit category name"
                  src={editImage}
                  onClick={() => this.toggleEditingCategoryName()} />
          </div>

          <div className={"edit-category-button " + this.editCategoryButtonShowHideClass()}>
            <img className="edit-category-image button-image"
                  alt="Edit this category"
                  src={moreImage}
                  onClick={this.toggleEditingCategory} />
          </div>

          <div className={"save-new-category-name-button " + this.saveNewCategoryNameButtonShowHideClass()}>
            <img className="save-new-category-name-image button-image"
                  alt="Save new categorty name"
                  src={checkmarkImage}
                  onClick={() => this.updateCategoryName(category.id)} />
          </div>

          <div className={"cancel-new-category-name-button " + this.cancelNewCategoryNameButtonShowHideClass()}>
            <img className="cancel-new-category-name-image button-image"
                  alt="Cance new category name"
                  src={deleteImage}
                  onClick={this.cancelCategoryNameUpdate} />
          </div>
        </div>

        <div>
          {
            this.state.showDialog &&
            <ReactConfirmAlert
              title="Confirm to submit"
              message="Are you sure to do this."
              confirmLabel="Confirm"
              cancelLabel="Cancel"
              onConfirm={() => alert('Action after Confirm')}
              onCancel={() => alert('Action after Cancel')}
            />
          }
        </div>
        <div className="category-items">
          {this.items()}
        </div>
        <div className="new-item-form">
          {this.itemForm()}
        </div>
      </div>
    );
  }
}

Category.propTypes = {
  category: PropTypes.object.isRequired,
  deleteCategoryHandler: PropTypes.func.isRequired,
  updateCategoryHandler: PropTypes.func.isRequired,
  list: PropTypes.object.isRequired,
  addItemHandler: PropTypes.func.isRequired,
  updateItemHandler: PropTypes.func.isRequired,
  deleteItemHandler: PropTypes.func.isRequired
};

export default Category;
