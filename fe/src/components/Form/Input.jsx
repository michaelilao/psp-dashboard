function Input({ field, state, setState }) {
	let value = state?.[field.id];
	if (field.input == "date") {
		let date = new Date();
		if (value) {
			date = new Date(value || "");
		}
		value = date.toISOString().split("T")[0];
	}
	return (
		<div key={field.id} className="mb-2">
			<label
				className="block text-sm font-medium text-gray-700 mb-1"
				htmlFor={field.id}
			>
				{field.label}
				{field.required ? "*" : ""}
			</label>
			<input
				required={field.required}
				type={field.input}
				onChange={(e) => {
					setState((curr) => {
						let newValue = e.target.value;
						if (field.input == "number") {
							newValue = Number(newValue);
						}
						return { ...curr, [field.id]: newValue };
					});
				}}
				className="w-full px-3 py-2 border border-gray-300 rounded-md text-sm"
				id={field.id}
				name={field.id}
				value={value}
			/>
		</div>
	);
}

export { Input };
