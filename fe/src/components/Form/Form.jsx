function Form({ state, setState, fields }) {
	return (
		<form onSubmit={(e) => e.preventDefault()}>
			{fields.map((field) => {
				if (field.input == "textarea") {
					return (
						<div key={field.id} className="mb-2">
							<label
								className="block text-sm font-medium text-gray-700 mb-1"
								htmlFor={field.id}
							>
								{field.label}
								{field.required ? "*" : ""}
							</label>
							<textarea
								required={field.required}
								type="text"
								onChange={(e) => {
									setState((curr) => {
										return { ...curr, [field.id]: e.target.value };
									});
								}}
								className="w-full px-3 py-2 border border-gray-300 rounded-md text-sm"
								id={field.id}
								name={field.id}
								value={state?.[field.id]}
							/>
						</div>
					);
				}

				if (field.input == "text" || field.input == "email") {
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
										return { ...curr, [field.id]: e.target.value };
									});
								}}
								className="w-full px-3 py-2 border border-gray-300 rounded-md text-sm"
								id={field.id}
								name={field.id}
								value={state?.[field.id]}
							/>
						</div>
					);
				}
			})}
		</form>
	);
}

export { Form };
