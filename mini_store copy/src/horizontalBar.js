import React, {useEffect, useState} from 'react';
import { Bar } from 'react-chartjs-2';

import {
    Chart as ChartJS,
    BarElement,
    CategoryScale,
    LinearScale,
    Tooltip,
    Legend
    
}from 'chart.js';


ChartJS.register( 
    BarElement,
    CategoryScale,
    LinearScale,
    Tooltip,
    Legend
);

function HorizontalBarChart() {
    const [chartData, setChartData] = useState(null);
  
    useEffect(() => {
      fetch('http://go-server:8000/products/') 
        .then(res => res.json())
        .then(data => {
          const labels = data.map(item => item.name);
          const quantities = data.map(item => item.quantity);
  
          setChartData({
            labels,
            datasets: [
              {
                label: 'Cantidad Vendida',
                data: quantities,
                backgroundColor: 'rgba(148, 106, 176, 0.6)',
              },
            ],
          });
        })
        .catch(err => console.error("Error al obtener datos:", err));
    }, []);
  
    if (!chartData) return <p>Cargando datos...</p>;
  
    return (
      <div>
        <Bar
          data={chartData}
          options={{
            indexAxis: 'y',
            responsive: true,
            plugins: {
              legend: {
                position: 'top',
              },
              title: {
                display: true,
                text: 'Top 10 productos vendidos en 2003',
              },
            },
          }}
        />
      </div>
    );
  }
  
export default HorizontalBarChart;


