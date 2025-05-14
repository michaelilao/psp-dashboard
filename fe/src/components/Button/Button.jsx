function Button({ onClick, label, isIcon = false, children }) {
	if (isIcon) {
		return (
			<button className="cursor-pointer" onClick={onClick}>
				{children}
			</button>
		);
	}

	return (
		<button
			className="px-3 py-1.5 bg-blue-300 text-blue-800 text-sm rounded hover:bg-blue-400 focus:outline-none focus:ring-2 focus:ring-blue-300 focus:ring-offset-1 transition"
			onClick={onClick}
		>
			{label}
		</button>
	);
}

export { Button };
