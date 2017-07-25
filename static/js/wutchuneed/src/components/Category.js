import React from "react";
import PropTypes from 'prop-types';

import Item from '../components/Item'

import moreImage from '../images/more-circle.png';
import editImage from '../images/edit-circle.svg';
import deleteImage from '../images/x-mark-circle.png';
import checkmarkImage from '../images/check-circle.svg';
import '../css/category.css';

class Category extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      showChildren: true,
      inEditMode: false,
      editingCategoryName: false,
      newCategoryName: ""
    }
  }

  toggleShowChildren = () => {
    console.log("toggle")
    this.setState({
      showChildren: !this.state.showChildren
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

  categoryEditFieldsShowHideClass = () => {
    if (this.state.inEditMode == true && this.state.editingCategoryName == false) {
      return "show"
    }
    return "hide"
  }

  categoryNameEditFieldsShowHideClass = () => {
    if (this.state.editingCategoryName == true) {
      return "show"
    }
    return "hide"
  }

  updateNewCategory = (e) => {
    this.setState({
      newCategoryName: e.target.value
    });
  }

  editCategoryButtonShowHideClass = () => {
    if (this.state.editingCategoryName == true) {
      return "hide"
    }
    return "show"
  }


  editCategoryInputShowHideClass = () => {
    if (this.state.editingCategoryName == true) {
      return "show"
    }
    return "hide"
  }

  categoryNameShowHideClass = () => {
    if (this.state.editingCategoryName == true) {
      return "hide"
    }
    return "show"
  }

  updateCategoryName = (categoryId) => {
    this.props.updateCategoryHandler(categoryId, this.state.newCategoryName)
    this.setState({
      newCategoryName: ""
    });
  }

  items() {
    if(this.state.showChildren) {
      return this.props.category.items.map((item, idx) => (
        <Item key={"item-" + item.id} item={item} />
      ));
    }
  }

  render() {
    const { category } = this.props;

    return (
      <div key={"cat-" + category.id} className={"category-row"}>
        <div className="category-header-row">
          <div className={"category-name " + this.categoryNameShowHideClass()} onClick={() => this.toggleShowChildren()}>
            {category.name}
          </div>
          <div className={"edit-category-name " + this.editCategoryInputShowHideClass()}>
            <input type="text"
                    id="edit-category-name-input"
                    value={this.state.newCategoryName}
                    onChange={this.updateNewCategory} />
          </div>
          <div className={"edit-category-button " + this.editCategoryButtonShowHideClass()}>
            <img className="edit-category-image button-image" src={moreImage} onClick={() => this.toggleEditingCategory()} />
          </div>
          <div className={"edit-category-buttons-container " + this.categoryEditFieldsShowHideClass()}>
            <div className="edit-category-name-button">
              <img className="edit-category-name-image button-image" src={editImage} onClick={() => this.toggleEditingCategoryName()} />
            </div>
            <div className="delete-category-button">
              <img className="delete-category-image button-image" src={deleteImage} onClick={() => this.props.deleteCategoryHandler(category.id)} />
            </div>
          </div>


          <div className={"edit-category-name-buttons-container " + this.categoryNameEditFieldsShowHideClass()}>
            <div className="edit-category-name-button">
              <img className="edit-category-name-image button-image" src={checkmarkImage} onClick={() => this.updateCategoryName(category.id)} />
            </div>
            <div className="delete-category-button">
              <img className="delete-category-image button-image" src={deleteImage} onClick={() => {} } />
            </div>
          </div>
        </div>
        <div className="category-items">
          {this.items()}
        </div>
      </div>
    );
  }
}

Category.propTypes = {
  category: PropTypes.object.isRequired,
  deleteCategoryHandler: PropTypes.func.isRequired,
  updateCategoryHandler: PropTypes.func.isRequired
};

export default Category;

//this.props.deleteCategoryHandler(category.id)
