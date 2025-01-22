export default class Task {
  #id;
  #projectID
  #title;
  #description;
  #endDate;
  #createdAt
  #category
  #status;
  constructor(id, projectID, title, description, endDate, createdAt, category, status = false) {
    this.#id = id;
    this.#projectID = projectID;
    this.#title = title;
    this.#description = description;
    this.#endDate = endDate;
    this.#createdAt = createdAt;
    this.#category = category;
    this.#status = status;
  }

  getID() {
    return this.#id;
  }
  getProjectID() {
    return this.#projectID;
  }
  getTitle() {
    return this.#title;
  }
  getDescription() {
    return this.#description;
  }
  getStatus() {
    return this.#status;
  }
  getEndDate() {
    return this.#endDate;
  }
  getCreatedAt() {
    return this.#createdAt;
  }
  getCategory() {
    return this.#category;
  }

  setID(id) {
    this.#id = id;
  }
  setProjectID(projectID) {
    this.#projectID = projectID;
  }
  setTitle(title) {
    this.#title = title;
  }
  setDescription(description) {
    this.#description = description;
  }
  setStatus(status) {
    this.#status = status;
  }
  setEndDate(endDate) {
    this.#endDate = endDate;
  }
  setCreatedAt(createdAt) {
    this.#createdAt = createdAt;
  }
  setCategory(category) {
    this.#category = category;
  }

  copyOf(task) {
    task.setID(this.#id);
    task.setProjectID(this.#projectID);
    task.setTitle(this.#title);
    task.setDescription(this.#description);
    task.setEndDate(this.#endDate);
    task.setCreatedAt(this.#createdAt);
    task.setCategory(this.#category);
    task.setStatus(this.#status);
  }
}
