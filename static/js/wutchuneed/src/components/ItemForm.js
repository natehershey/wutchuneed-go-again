import React from "react";
import PropTypes from 'prop-types';

import '../css/itemForm.css';

class ItemForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
        itemName : "",
        itemQuantity : 0,
        itemMeasure : "",
        itemStatus : "needed"
    };
  }

  handleInputChange = (event) => {
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    this.setState({
      [name]: value
    });
  }

  handleSubmit = (event) => {
    event.preventDefault();

    var item = {
      name: this.state.itemName,
      quantity: parseInt(this.state.itemQuantity, 10),
      measure: this.state.itemMeasure,
      status: this.state.itemStatus,
      listId: this.props.list.id,
      categoryId: this.props.category.id
    }

    console.log("About to call addItemHandler on submit")
    console.log("payload: ", item)
    this.props.addItemHandler(item)
    this.props.closeMe()

  }

  render() {
    return(
      <div>
        <div className="close-add-item-dialog" onClick={this.props.handleClose}>
          &times;
        </div>
        <div className="add-item-form-header">
          {`Add item to ${this.props.category.name}`}
        </div>

        <div className="add-item-form">
          <form onSubmit={this.handleSubmit}>
            <div className="item-name-field item-input-field">
              <div className="label">
                Name
              </div>
              <input name="itemName" className="item-name-input item-input" type="text" onChange={this.handleInputChange} />
            </div>
            <div className="item-quantity-field item-input-field">
              <div className="label">
                Quantity
              </div>
              <input name="itemQuantity" className="item-quantity-input item-input" type="text" onChange={this.handleInputChange} />
            </div>
            <div className="item-measure-field item-input-field">
              <div className="label">
                Measure
              </div>
              <input name="itemMeasure" className="item-measure-input item-input" type="text" onChange={this.handleInputChange} />
            </div>
            <input className="add-item-form-submit-button" type="submit" value="Submit" />
          </form>
        </div>
      </div>
    )
  }
}


ItemForm.propTypes = {
  category: PropTypes.object.isRequired,
  list: PropTypes.object.isRequired,
  addItemHandler: PropTypes.func.isRequired,
  closeMe: PropTypes.func.isRequired
};

export default ItemForm;
