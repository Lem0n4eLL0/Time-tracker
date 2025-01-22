export default class Project {
  #id;
  #title;
  #description;
  #createdTime
  #tasks;

  constructor(id, title, description, createdTime, tasks = new Map()) {
    this.#id = id;
    this.#title = title;
    this.#description = description;
    this.#tasks = tasks;
  }

  getID() {
    return this.#id;
  }
  getTitle() {
    return this.#title;
  }
  getDescription() {
    return this.#description;
  }
  getTasks() {
    return this.#tasks;
  }
  getTask(key) {
    return this.#tasks.get(key);
  }
  getCreatedTime() {
    return this.#createdTime;
  }

  setTitle(title) {
    this.#title = title;
  }
  setDescription(description) {
    this.#description = description;
  }
  setTasks(tasks) {
    this.#tasks = tasks;
  }
  setTask(key, value) {
    this.#tasks.set(key, value);
  }

  addTask(key, task) {
    return this.#tasks.set(key, task)
  }
  deleteTask(key) {
    return this.#tasks.delete(key)
  }
}


