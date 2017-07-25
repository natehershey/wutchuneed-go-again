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

  constructor(props) {
    super(props)
  }

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
        <div className="new-category">
          <form onSubmit={this.handleNewCategorySubmit}>
            <input type="text"
                    id="new-category-input-field"
                    value={this.state.newCategoryName}
                    onChange={(e) => this.updateNewCategory(e)} />
            <img className="new-category-save-image"
                  src={save}
                  onClick={() => this.handleNewCategorySubmit()}/>
            <img className="new-category-cancel-image"
                  src={cancel}
                  onClick={() => this.clearCategoryInput()}/>
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
                updateCategoryHandler={this.props.updateCategoryHandler}/>
    ));
  }

  render() {
    const { list } = this.props;

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
  updateCategoryHandler: PropTypes.func.isRequired
};

export default List;
