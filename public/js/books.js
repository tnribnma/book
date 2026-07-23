let allBooks = [];

async function loadBooks() {
  const res = await apiRequest('/books');
  if (res && res.ok) {
    const data = await res.json();
    allBooks = data.data || [];
    renderBooks(allBooks);
  }
}

function renderBooks(books) {
  const tbody = document.getElementById('booksBody');
  tbody.innerHTML = '';

  if (books.length === 0) {
    tbody.innerHTML = `<tr><td colspan="4" class="text-center py-8 text-gray-500">No books found</td></tr>`;
    return;
  }

  books.forEach(book => {
    const row = document.createElement('tr');
    row.innerHTML = `
      <td class="font-medium">${book.title}</td>
      <td>${book.author}</td>
      <td>${book.published_year || '-'}</td>
      <td>
        <button onclick="editBook(${book.id})" 
                class="text-blue-600 hover:text-blue-800 mr-3">Edit</button>
        <button onclick="deleteBook(${book.id})" 
                class="text-red-600 hover:text-red-800">Delete</button>
      </td>
    `;
    tbody.appendChild(row);
  });
}

function filterBooks() {
  const query = document.getElementById('searchInput').value.toLowerCase();
  const filtered = allBooks.filter(book => 
    book.title.toLowerCase().includes(query) || 
    book.author.toLowerCase().includes(query)
  );
  renderBooks(filtered);
}

async function addBook() {
  const title = document.getElementById('title').value.trim();
  const author = document.getElementById('author').value.trim();
  const year = document.getElementById('year').value;

  if (!title || !author) {
    alert("Title and Author are required!");
    return;
  }

  const res = await apiRequest('/books', 'POST', {
    title, author, published_year: year ? parseInt(year) : null
  });

  if (res && res.ok) {
    loadBooks();
    document.getElementById('title').value = '';
    document.getElementById('author').value = '';
    document.getElementById('year').value = '';
  }
}

async function deleteBook(id) {
  if (confirm("Are you sure you want to delete this book?")) {
    const res = await apiRequest(`/books/${id}`, 'DELETE');
    if (res && res.ok) loadBooks();
  }
}

function editBook(id) {
  alert("Edit feature coming soon! (We can add modal later)");
  // You can extend this with a modal for editing
}

function logout() {
  localStorage.removeItem('token');
  window.location.href = 'login.html';
}

// Load on page start
window.onload = loadBooks;