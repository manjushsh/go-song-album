const apiUrl = 'https://go-song-album.onrender.com';
let authToken = localStorage.getItem('authToken');

document.addEventListener('DOMContentLoaded', () => {
    if (authToken) fetchAlbums();
});

async function login() {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    try {
        const data = await fetchData(`${apiUrl}/auth/login`, 'POST', { username, password });
        authToken = data.token;
        localStorage.setItem('authToken', authToken);
        fetchAlbums();
    } catch (error) {
        alert(error.message);
    }
}

async function fetchAlbums() {
    try {
        const albums = await fetchData(`${apiUrl}/albums`);
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
                <button onclick="deleteAlbum(${album.id})">Delete</button>
            </div>
        </div>
    `).join('');
    toggleView('album-manager');
}

async function addAlbum() {
    const album = getAlbumFormData();

    try {
        await fetchData(`${apiUrl}/albums`, 'POST', album);
        fetchAlbums().then(() => setAlbumFormData({ title: '', artist: '', price: null, image: '' }));
    } catch (error) {
        alert(error.message);
    }
}

async function updateAlbum() {
    const album = getAlbumFormData();
    const id = document.getElementById('albumId').value;

    try {
        await fetchData(`${apiUrl}/albums/${id}`, 'PUT', album);
        fetchAlbums().then(() => setAlbumFormData({ title: '', artist: '', price: null, image: '' }));
    } catch (error) {
        alert(error.message);
    }
}

async function deleteAlbum(id) {
    try {
        await fetchData(`${apiUrl}/albums/${id}`, 'DELETE');
        fetchAlbums();
    } catch (error) {
        alert(error.message);
    }
}

async function resetAlbums() {
    try {
        await fetchData(`${apiUrl}/albums/reset`, 'POST');
        fetchAlbums();
    } catch (error) {
        alert(error.message);
    }
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
    localStorage.removeItem('authToken');
    window.location.reload();
}

function toggleView(viewId) {
    document.getElementById('login-form').style.display = viewId === 'login-form' ? 'block' : 'none';
    document.getElementById('album-manager').style.display = viewId === 'album-manager' ? 'block' : 'none';
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
