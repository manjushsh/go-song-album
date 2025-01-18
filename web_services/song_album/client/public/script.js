const apiUrl = 'https://go-song-album.onrender.com';
let authToken = localStorage.getItem('authToken');
let isLoading = false;

document.addEventListener('DOMContentLoaded', () => {
    if (authToken) fetchAlbums();
});

async function login() {
    const { username, password } = getFormData('username', 'password');
    if (!username || !password) return alert('Please enter a username and password');

    try {
        isLoading = true;
        const data = await fetchData(`${apiUrl}/auth/login`, 'POST', { username, password });
        authToken = data.token;
        localStorage.setItem('authToken', authToken);
        fetchAlbums().finally(() => isLoading = false);
    } catch (error) {
        alert(error.message);
    }
}

async function register() {
    const { 'register-username': username, 'register-password': password, 'register-confirm-password': rpassword } = getFormData('register-username', 'register-password', 'register-confirm-password');
    if (!username || !password || !rpassword) return alert('Please enter a username and password');

    try {
        await fetchData(`${apiUrl}/auth/register`, 'POST', { username, password });
        alert('User registered successfully');
    } catch (error) {
        alert(error.message);
    }
}

async function fetchAlbums() {
    try {
        isLoading = true;
        const albums = await fetchData(`${apiUrl}/albums`)
            .finally(() => isLoading = false);
        renderAlbums(albums);
    } catch (error) {
        alert(error.message);
    }
}

function renderAlbums(albums) {
    const albumsDiv = document.getElementById('albums');
    albumsDiv.innerHTML = albums.map(album => `
        <div class="album" id="${album.id}">
            <img src="${album.image}" alt="${album.title}">
            <div class="album-details">
                <h3>${album.title}</h3>
                <p>${album.artist}</p>
                <p>$${album.price}</p>
            </div>
            <div class="album-actions">
                <button data-id="${album.id}" onclick="editAlbum(this)">Edit</button>
                <button data-id="${album.id}" onclick="deleteAlbum(this)">Delete</button>
            </div>
        </div>
    `).join('');
    toggleView('album-manager');
}

async function addAlbum() {
    const album = getAlbumFormData();

    try {
        isLoading = true;
        await fetchData(`${apiUrl}/albums`, 'POST', album);
        await fetchAlbums();
        setAlbumFormData({ title: '', artist: '', price: null, image: '' });
    } catch (error) {
        alert(error.message);
    }
}

async function updateAlbum() {
    isLoading = true;
    const album = getAlbumFormData();
    const id = document.getElementById('albumId').value;

    try {
        await fetchData(`${apiUrl}/albums/${id}`, 'PUT', album);
        await fetchAlbums();
        setAlbumFormData({ title: '', artist: '', price: null, image: '' });
    } catch (error) {
        alert(error.message);
    }
}

async function deleteAlbum(button) {
    const id = button.getAttribute('data-id');
    try {
        isLoading = true;
        await fetchData(`${apiUrl}/albums/${id}`, 'DELETE');
        fetchAlbums().finally(() => isLoading = false);
    } catch (error) {
        alert(error.message);
    }
}

async function resetAlbums() {
    window.location.reload();
}

async function editAlbum(button) {
    const id = button.getAttribute('data-id');
    try {
        const album = await fetchData(`${apiUrl}/albums/${id}`);
        setAlbumFormData(album);
    } catch (error) {
        alert(error.message);
    }
}

function logout() {
    localStorage.removeItem('authToken');
    window.location.reload();
}

function handleUnauthorized() {
    toggleView('login-form');
    toggleView('register-form');
    localStorage.removeItem('authToken');
    window.location.reload();
}

function toggleView(viewId) {
    ['login-form', 'register-form', 'album-manager'].forEach(id => {
        document.getElementById(id).style.display = id === viewId ? 'block' : 'none';
    });
}

function getAuthHeaders() {
    return {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authToken}`
    };
}

function getAlbumFormData() {
    return {
        id: document.getElementById('albumId').value ?? null,
        title: document.getElementById('title').value,
        artist: document.getElementById('artist').value,
        price: parseFloat(document.getElementById('price').value) ?? 10.69,
        image: document.getElementById('imageUrl').value
    };
}

function setAlbumFormData(album) {
    document.getElementById('albumId').value = album.id;
    document.getElementById('title').value = album.title;
    document.getElementById('artist').value = album.artist;
    document.getElementById('price').value = album.price;
    document.getElementById('imageUrl').value = album.image;
}

function getFormData(...ids) {
    return ids.reduce((data, id) => {
        data[id] = document.getElementById(id).value;
        return data;
    }, {});
}

async function fetchData(url, method = 'GET', body = null) {
    const options = {
        method,
        headers: getAuthHeaders()
    };
    if (body) options.body = JSON.stringify(body);

    const response = await fetch(url, options);

    if (response.status === 401) {
        handleUnauthorized();
        throw new Error('Unauthorized');
    }

    if (!response.ok) throw new Error('Request failed');

    return response.json();
}
