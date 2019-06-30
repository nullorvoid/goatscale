class WebApiClient {
	// Handle the error codes and passing the data into the async promise
	private handleRequestResponse<T>(req: XMLHttpRequest, resolve: (value?: T | PromiseLike<T>) => void, reject: (reason?: any) => void) {
		req.onload = () => {
			// If we are not 200 explode wildly
			if (req.status !== 200) {
				return reject('request returned a non 200 status code: ' + req.status);
			}

			// If we do not have any response text resolve
			if (!req.responseText){
				return resolve();
			}

			// If we have response text parse to json, if not valid explode
			try {
				const resData = JSON.parse(req.responseText);
				return resolve(resData);
			} catch (err) {
				return reject(err);
			}
		};

		// If the request hard errors, explode some more
		req.onerror = () => {
			return reject('request error');
		};
	}

	// post a request to the server
	private post(path: string, data: any): Promise<any> {
		return new Promise((resolve, reject) => {
			const req = new XMLHttpRequest();

			this.handleRequestResponse<any>(req, resolve, reject);

			req.open('POST', path);
			req.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
			req.send(JSON.stringify(data));
		});
	}

	// Used to create a session for the user which can be used as a valid request to any server
	public login(userId: string): Promise<any> {
		const query = {
			userId: userId,
		};

		return this.post('/login', query);
	}
}

// Export this directly as we do not have any config and only need one instance
export const webApiClient = new WebApiClient();
