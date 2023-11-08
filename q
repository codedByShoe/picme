{{define "content"}}
<div class="py-12 flex justify-center">
	<div class="px-8 py-8 bg-white shadow rounded">
		<h1 class="pt-4 pb-8 text-center text-3xl font-bold text-gray-900">Start Sharing Your Photos Today</h1>
		<form action="/users" method="POST">
			<div class="py-2">
				<label for="email" class="text-sm font-semibold text-gray-800">Email Address</label>
				<input type="email" name="email" id="email" placeholder="Your Email"
					class="w-full px-3 py-2 border border-gray-300 placeholder-gray-300 text-gray-800 rounded"
					autocomplete="off" />
			</div>
			<div class="py-2">
				<label for="password" class="text-sm font-semibold text-gray-800">Password</label>
				<input type="password" id="password" name="password" placeholder="Password"
					class="w-full px-3 py-2 border border-gray-300 placeholder-gray-300 text-gray-800 rounded"
					autocomplete="off" />
			</div>
			<div class="py-4">
				<button class="w-full py-4 px-2 bg-sky-600 hover:bg-sky-700 rounded font-bold text-lg text-white"
					type="submit">Sign Up</button>
			</div>
			<div class="py-2 w-full flex justify-between">
				<p>Already have an account? <a href="/signin">Sign in</a></p>
				<p><a href="/reset-password">Forgot your password?</a></p>
			</div>
		</form>
	</div>
</div>
{{end}}
