const sleep = (ms) => new Promise(res => setTimeout(res, ms));

const typewriter = async (text, selector, delay) => {
	for (let i = 0; i < text.length; i++) {
		document.querySelector(selector).textContent += text[i];
		await sleep(delay);
	}
}
