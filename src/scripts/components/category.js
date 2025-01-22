export default class Category {
  #id;
  #name;
  #description;

  constructor(id, name, description) {
    this.#id = id;
    this.#name = name;
    this.#description = description;
  }

  getID() {
    return this.#id;
  }
  getName() {
    return this.#name;
  }
  getDescription() {
    return this.#description;
  }

  setName(name) {
    this.#name = name;
  }
  setDescription(description) {
    this.#description = description;
  }

  createOption() {
    const element = document.createElement('option');
    element.classList.add('category__option');
    element.value = this.#id;
    element.textContent = this.#name;
    return element;
  }
}
