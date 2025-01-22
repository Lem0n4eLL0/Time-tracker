import * as general from './common'

function loginRequest(message) {
    fetch(general.createHost('api/user/login'), {
      method: 'POST',
      body: JSON.stringify({
        username: message.username,
        password: message.password
      }),
      headers: {
          "Content-Type": "application/json",
      }
    })
    .then((res) => {
      if(res.ok) {
        return ;
      }
      return Promise.reject(res.status);
    })
    .then(() => {
      window.location.assign(general.createHost('profile'));
    })
    .catch(err => console.log(`Ошибка: ${err}`));
}

function registrationRequest(message) {
    fetch(general.createHost('api/user/registration'), {
    method: 'POST',
    body: JSON.stringify({
	    username: message.username,
	    password: message.password,
      email: message.email,
    }),
    headers: {
        "Content-Type": "application/json",
    }
    })
    .then((res) => {
        if(res.ok) {
            return ;
        }
        return Promise.reject(res.status);
    })
    .then(() => {
      location.reload(true)
    })
    .catch(err => console.log(`Ошибка: ${err}`));
}

export function getProjects() {
  return fetch(general.createHost('api/user/projects'), {
      method: 'GET',
      headers: {
          "Content-Type": "application/json",
      }
    })
}

export function addProject(message) {
  return fetch(general.createHost('api/user/projects'), {
    method: 'POST',
    headers: {
        "Content-Type": "application/json",
    },
    body: JSON.stringify({
      project_name: message.name,
	    project_description: message.description,
    })
  })
}

export function deleteProject(projectID) {
  return fetch(general.createHost(`api/user/projects/${projectID}`), {
    method: 'DELETE',
    headers: {
        "Content-Type": "application/json",
    },
  })
}

export function changeProject(message) {
  return fetch(general.createHost(`api/user/projects/${message.id}`), {
    method: 'PUT',
    headers: {
        "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id: message.id,
      project_name: message.name,
      project_description: message.description
    })
  })
}

export function getTasks(projectID) {
  return fetch(general.createHost(`api/projects/${projectID}/tasks`), {
      method: 'GET',
      headers: {
          "Content-Type": "application/json",
      }
    })
}

export function addTask(message) {
  return fetch(general.createHost(`api/projects/${message.projectID}/tasks`), {
    method: 'POST',
    headers: {
        "Content-Type": "application/json",
    },
    body: JSON.stringify({
      projectID: message.projectID,
      task_name: message.task_name,
      task_description: message.task_description,
      status: false,
      category_name: message.category_name
    })
  })
}

export function getTaskCategories() {
  return fetch(general.createHost(`api/user/tasks/categories`), {
    method: 'GET',
    headers: {
        "Content-Type": "application/json",
    },
  })
}

export function deleteTask(projectID, taskID) {
  return fetch(general.createHost(`api/projects/${projectID}/tasks/${taskID}`), {
    method: 'DELETE',
    headers: {
        "Content-Type": "application/json",
    },
  })
}

export function changeTask(task) {
  return fetch(general.createHost(`api/projects/${task.getProjectID()}/tasks/${task.getID()}`), {
    method: 'PUT',
    headers: {
        "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id: task.getID(),
      projectID: task.getProjectID(),
      task_name: task.getTitle(),
      task_description: task.getDescription(),
      status: task.getStatus(),
      end_date: task.getEndDate(),
      category_name: task.getCategory()
    })
  })
}

export function getReportPDF() {
  return fetch(general.createHost('api/projects/report/pdf'), {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });
}


 export {loginRequest as login, registrationRequest as registration, getReportPDF as getPDF}
// ,getProjects, addProject, deleteProject, changeProject, getTasks, addTask, deleteTask}
