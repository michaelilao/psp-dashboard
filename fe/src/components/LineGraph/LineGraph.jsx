import {
	Chart as ChartJS,
	CategoryScale,
	LinearScale,
	PointElement,
	LineElement,
	Title,
	Tooltip,
	Legend,
} from "chart.js";
import { Line } from "react-chartjs-2";

ChartJS.register(
	CategoryScale,
	LinearScale,
	PointElement,
	LineElement,
	Title,
	Tooltip,
	Legend
);

export function LineGraph({ data, title }) {
	return (
		<Line
			data={data}
			options={{
				responsive: true,
				maintainAspectRatio: false,
				layout: {
					padding: {
						bottom: 0,
					},
				},
				plugins: {
					title: {
						display: true,
						text: title,
						align: "center",
					},
					legend: {
						display: false,
						position: "left",
					},
				},
			}}
		/>
	);
}
