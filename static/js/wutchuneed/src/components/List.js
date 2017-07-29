import React from "react";
import PropTypes from 'prop-types';

import Category from '../components/Category'

import save from '../images/check-circle.svg';
import cancel from '../images/x-mark-circle.png';
import '../css/list.css';

class List extends React.Component {
  state = {
    addingCategory: true,
    newCategoryName: ""
  };

  updateNewCategory = (e) => {
    this.setState({
      newCategoryName: e.target.value
    });
  }

  handleNewCategorySubmit(event) {
    const newCategoryName = this.state.newCategoryName;
    const listId = this.props.list.id;

    this.props.addCategoryHandler(newCategoryName, listId);
    this.clearCategoryInput();
  }

  clearCategoryInput() {
    this.setState({
      newCategoryName: ""
    });
  }

  newCategory() {
    if (this.state.addingCategory) {
      return(
        <div>
          <form onSubmit={this.handleNewCategorySubmit}>
            <div className="new-category">
              <div className="new-category-input-container">
                <input type="text"
                        id="new-category-input-field"
                        value={this.state.newCategoryName}
                        onChange={(e) => this.updateNewCategory(e)} />
              </div>
              <div className="new-catgegory-save-button">
                <img className="new-category-save-image"
                      alt="Save new category"
                      src={save}
                      onClick={() => this.handleNewCategorySubmit()}/>
              </div>
              <div className="new-category-cancel-button">
                <img className="new-category-cancel-image"
                      alt="Cancel new category"
                      src={cancel}
                      onClick={() => this.clearCategoryInput()}/>
              </div>
            </div>
          </form>
        </div>
      );
    }
  }

  categoryRows() {
    return this.props.list.categories.map((category, idx) => (
      <Category key={"category-" + category.id}
                category={category}
                deleteCategoryHandler={this.props.deleteCategoryHandler}
                updateCategoryHandler={this.props.updateCategoryHandler}
                list={this.props.list}
                addItemHandler={this.props.addItemHandler}
                updateItemHandler={this.props.updateItemHandler}
                deleteItemHandler={this.props.deleteItemHandler}
                />
    ));
  }

  render() {
    return (
      <div className={"list-detail-view"}>
        {this.newCategory()}
        <div className="list-category">
          {this.categoryRows()}
        </div>
      </div>
    );
  }
}

List.propTypes = {
  list: PropTypes.object.isRequired,
  addCategoryHandler: PropTypes.func.isRequired,
  deleteCategoryHandler: PropTypes.func.isRequired,
  updateCategoryHandler: PropTypes.func.isRequired,
  addItemHandler: PropTypes.func.isRequired,
  updateItemHandler: PropTypes.func.isRequired,
  deleteItemHandler: PropTypes.func.isRequired
};

export default List;
