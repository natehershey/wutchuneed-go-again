function getLists(cb) {
  return fetch("api/v1/lists", {
    accept: "application/json"
  })
    .then(checkStatus)
    .then(parseJSON)
    .then(cb);
}

function getList(id, cb) {
  return fetch(`api/v1/lists/${id}`, {
    accept: "application/json"
  })
    .then(checkStatus)
    .then(parseJSON)
    .then(cb);
}

function addCategory(name, listId, cb) {
  return fetch("api/v1/categories", {
    method: "POST",
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      listId: listId,
      name: name
    })
  })
  .then(checkStatus)
  .then(parseJSON)
  .then(cb);
}

function deleteCategory(id, cb) {
  return fetch(`api/v1/categories/${id}`, {
    method: "DELETE",
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
    }
  })
  .then(checkStatus)
  .then(parseJSON)
  .then(cb);
}

function updateCategory(id, attributes) {
  // TODO: Implement in API and here
  console.log("updateCategoryName(id, name) NOT IMPLEMENTED");

  // return fetch(`api/v1/categories/${id}`, {
  //   method: "PUT",
  //   headers: {
  //     'Accept': 'application/json',
  //     'Content-Type': 'application/json',
  //   },
  //   body: {
  //     name: name
  //   }
  // })
  // .then(checkStatus)
  // .then(parseJSON)
  // .then(cb);

}

function checkStatus(response) {
  if (response.status >= 200 && response.status < 300) {
    return response;
  }
  const error = new Error(`HTTP Error ${response.statusText}`);
  error.status = response.statusText;
  error.response = response;
  console.log(error); // eslint-disable-line no-console
  throw error;
}

function parseJSON(response) {
  return response.json();
}

const Client = { getLists, getList, addCategory, deleteCategory, updateCategory };

export default Client;
