import { Chart as ChartJS, ArcElement, Tooltip, Legend, Title } from "chart.js";
import { Pie } from "react-chartjs-2";

const colors = [
	"rgba(255, 99, 132)",
	"rgba(54, 162, 235)",
	"rgba(255, 206, 86)",
	"rgba(75, 192, 192)",
	"rgba(153, 102, 255)",
	"rgba(255, 159, 64)",
];

ChartJS.register(ArcElement, Tooltip, Legend, Title);

function PieChart({ data, title }) {
	return (
		<Pie
			data={data}
			options={{
				plugins: {
					title: {
						display: true,
						text: title,
						align: "center",
					},
					legend: {
						display: true,
						position: "left",
					},
				},
			}}
		/>
	);
}

export { PieChart, colors };
