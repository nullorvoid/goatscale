import { webApiClient } from "./webapiclient";

// Function to generate a user id, this is a simple version of a UUID
function generateGuid() {
	return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, (c) => {
		const r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
		return v.toString(16);
	});
}

// Get user id retrieves the user id from localstorage, if it is not there is will create it
function getUserId(): string {
	let userId = window.localStorage.getItem('userId');

	if (!userId) {
		userId = generateGuid();
		window.localStorage.setItem('userId', userId);
	}

	return userId;
}

// create a login button on the page so we can demonstrate the login route
function createLoginBtn() {
	const button = document.createElement('button');
	button.innerText = 'Login';
	button.addEventListener("click", (e: Event) => {
		const userId = getUserId();
		webApiClient.login(userId)
			.then(() => { console.log('Logged in with user id: ' + userId); })
			.catch((error) => { console.log('Login Failed: ' + error); });
	});

	return button;
}

// Add login button to the body
document.body.appendChild(createLoginBtn());
