import { TextArea } from "./Textarea";
import { Input } from "./Input";
import { Select } from "./Select";

function Form({ state, setState, fields }) {
	return (
		<form onSubmit={(e) => e.preventDefault()}>
			{fields.map((field) => {
				if (field.input == "textarea") {
					return <TextArea field={field} state={state} setState={setState} />;
				}
				if (field.input == "select") {
					return <Select field={field} state={state} setState={setState} />;
				}

				return <Input field={field} state={state} setState={setState} />;
			})}
		</form>
	);
}

export { Form };
