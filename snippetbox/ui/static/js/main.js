var navLinks = document.querySelectorAll("nav a");
for (var i = 0; i < navLinks.length; i++) {
	var link = navLinks[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}

document.addEventListener("DOMContentLoaded", function() {
	document.getElementById("login-button").addEventListener("click", function() {
		// Отправить запрос на страницу входа
		window.location.href = "/user/login";
	});

	document.getElementById("signup-button").addEventListener("click", function() {
		// Отправить запрос на страницу регистрации
		window.location.href = "/user/signup";
	});
});
