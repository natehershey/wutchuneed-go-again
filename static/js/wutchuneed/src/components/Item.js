import React from "react";
import PropTypes from 'prop-types';

import deleteImage from '../images/x-mark-circle.png';
import cartImage from '../images/cart-circle.svg';

import '../css/item.css';

class Item extends React.Component {
  addItemToCart = () => {
    var updated_item
    if (this.props.item.status === 'in_cart') {
      updated_item = {
        "status": "needed"
      }
      this.props.updateItemHandler(this.props.item.id, updated_item)
    } else {
      updated_item = {
        "status": "in_cart"
      }
      this.props.updateItemHandler(this.props.item.id, updated_item)
    }
  }

  strikethroughClass() {
    if (this.props.item.status === "in_cart") {
      return "in-cart"
    }
    return null
  }

  render() {
    const { item } = this.props;

    return (
      <div key={"item-" + item.id} className="item-row">
        <div className="left-side">
          <div className="item-delete-button">
            <img className="item-delete-image"
                  alt="Delete Item"
                  src={deleteImage}
                  onClick={() => {this.props.deleteItemHandler(item.id)}} />
          </div>
          <div className={"item-name " + this.strikethroughClass()}
                onClick={this.addItemToCart} >
            {item.name}
          </div>
        </div>
        <div className="right-side">
          <div className="add-to-cart-button">
            <img src={cartImage}
                  alt="Add to cart"
                  className="add-to-cart-image"
                  onClick={this.addItemToCart} />
          </div>
        </div>
      </div>
    );
  }
}

Item.propTypes = {
  item: PropTypes.object.isRequired,
  deleteItemHandler: PropTypes.func.isRequired
};

export default Item;
