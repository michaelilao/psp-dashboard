import { TextArea } from "./Textarea";
import { Input } from "./Input";
import { Select } from "./Select";

function Form({ state, setState, fields }) {
	return (
		<form onSubmit={(e) => e.preventDefault()}>
			{fields.map((field) => {
				const props = {
					field: field,
					state: state,
					setState: setState,
				};

				if (field.input == "textarea") {
					return <TextArea key={field.id} {...props} />;
				}
				if (field.input == "select") {
					return <Select key={field.id} {...props} />;
				}
				return <Input key={field.id} {...props} />;
			})}
		</form>
	);
}

export { Form };
