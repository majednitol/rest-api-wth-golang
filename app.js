// API URLs
const API_BASE_URL = 'http://localhost:3000/api';
const LOGIN_URL = 'http://localhost:3000/login';
const REGISTER_URL = 'http://localhost:3000/register';

// Helper to set token in localStorage
function setToken(token) {
  localStorage.setItem('token', token);
}

// Helper to get token from localStorage
function getToken() {
  return localStorage.getItem('token');
}

// Helper to show sections
function showSection(sectionId) {
  document.querySelectorAll('.section').forEach((section) => {
    section.style.display = 'none';
  });
  document.getElementById(sectionId).style.display = 'block';
}

// Login function
document.getElementById('login-form').addEventListener('submit', async (event) => {
  event.preventDefault();
  const username = document.getElementById('login-username').value;
  const password = document.getElementById('login-password').value;

  try {
    const response = await fetch(LOGIN_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });
    if (!response.ok) {
      throw new Error('Invalid credentials');
    }
    const data = await response.json();
    setToken(data.token);
    showSection('task-section');
    fetchTasks(); // Fetch tasks after login
  } catch (error) {
    alert(error.message);
  }
});

// Register function
document.getElementById('register-form').addEventListener('submit', async (event) => {
  event.preventDefault();
  const username = document.getElementById('register-username').value;
  const password = document.getElementById('register-password').value;

  try {
    const response = await fetch(REGISTER_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });
    if (!response.ok) {
      throw new Error('Registration failed');
    }
    alert('User registered successfully');
    showSection('login-section');
  } catch (error) {
    alert(error.message);
  }
});

// Task functions
document.getElementById('create-task-form').addEventListener('submit', async (event) => {
  event.preventDefault();
  const title = document.getElementById('task-title').value;
  const detail = document.getElementById('task-detail').value;

  try {
    const response = await fetch(`${API_BASE_URL}/tasks`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${getToken()}`,
      },
      body: JSON.stringify({ title, detail }),
    });
    if (!response.ok) {
      throw new Error('Failed to create task');
    }
    fetchTasks(); // Refresh task list
  } catch (error) {
    alert(error.message);
  }
});

// Update task
document.getElementById('update-task-form').addEventListener('submit', async (event) => {
  event.preventDefault();
  const id = document.getElementById('update-task-id').value;
  const title = document.getElementById('update-task-title').value;
  const detail = document.getElementById('update-task-detail').value;

  try {
    const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${getToken()}`,
      },
      body: JSON.stringify({ title, detail }),
    });
    if (!response.ok) {
      throw new Error('Failed to update task');
    }
    fetchTasks(); // Refresh task list
  } catch (error) {
    alert(error.message);
  }
});

// Fetch tasks
async function fetchTasks() {
  try {
    const response = await fetch(`${API_BASE_URL}/tasks`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${getToken()}`,
      },
    });
    if (!response.ok) {
      throw new Error('Failed to fetch tasks');
    }
    const tasks = await response.json();
    const taskList = document.getElementById('task-list');
    taskList.innerHTML = '';
    tasks.forEach((task) => {
      const taskElement = document.createElement('div');
      taskElement.classList.add('task-item');
      taskElement.innerHTML = `
        <strong>${task.title}</strong> <br> ${task.detail}
        <button onclick="deleteTask('${task.id}')">Delete</button>
      `;
      taskList.appendChild(taskElement);
    });
  } catch (error) {
    alert(error.message);
  }
}

// Delete task
async function deleteTask(id) {
  try {
    const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${getToken()}`,
      },
    });
    if (!response.ok) {
      throw new Error('Failed to delete task');
    }
    fetchTasks(); // Refresh task list
  } catch (error) {
    alert(error.message);
  }
}

// Show/hide sections
document.getElementById('show-register').addEventListener('click', () => {
  showSection('register-section');
});
document.getElementById('show-login').addEventListener('click', () => {
  showSection('login-section');
});

// Initially, show login section
showSection('login-section');
