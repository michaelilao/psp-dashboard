const categories = [
  "Rent & Utilities",
  "Groceries",
  "Transportation",
  "Dining & Entertainment",
  "Healthcare",
  "Shopping",
  "Salary",
  "Freelance Income",
  "Investments",
  "Savings & Loans",
  'Other'
];


const FORM_FIELDS = [
	{ id: "name", label: "Name", input: "text" },
	{ id: "amount", label: "Amount", input: "number", required: true },
	{
		id: "category",
		label: "Category",
		input: "select",
		options: categories,
		required: true,
	},
	{
		id: "transactionType",
		label: "Type",
		input: "select",
		options: ["income", "expense"],
		required: true,
	},
	{ id: "date", label: "Date", input: "date", required: true },
	{ id: "notes", label: "Notes", input: "textarea" },
];

export { FORM_FIELDS }