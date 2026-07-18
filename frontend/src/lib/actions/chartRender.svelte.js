import { Chart } from 'chart.js/auto';

export const chartRender = (node, options) => {
  let chart;

  const rafId = requestAnimationFrame(() => {
    chart = new Chart(node, {
      ...options,
      options: {
        ...options.options,
        responsive: true,
        maintainAspectRatio: false,
        layout: {
          padding: {
            top: 20,
            bottom: 20,
            left: 10,
            right: 10
          }
        }
	  }
     });
  });

  return {
    update(newOptions) {
      if (chart) {
        chart.data = newOptions.data;
        chart.update();
      }
    },
    destroy() {
      cancelAnimationFrame(rafId);
      chart?.destroy();
    }
  };
};
