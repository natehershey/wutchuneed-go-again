import React from "react";
import PropTypes from 'prop-types';

import '../css/item.css';

class Item extends React.Component {
  constructor(props) {
    super(props)
  }

  render() {
    const { item } = this.props;

    return (
      <div key={"item-" + item.id} className={"item-row"}>
        <div className="item-name">
          {item.name}
        </div>
      </div>
    );
  }
}

Item.propTypes = {
  item: PropTypes.object.isRequired
};

export default Item;
