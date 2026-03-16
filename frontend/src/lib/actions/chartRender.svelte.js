import { Chart } from 'chart.js/auto';

export const chartRender = (node, options) => {
  const chart = new Chart(node, {
    ...options,
    options: {
      ...options.options,
      responsive: true,
      maintainAspectRatio: false,
    }
  });

  return {
    update(newOptions) {
      chart.data = newOptions.data;
      chart.update();
    },
    destroy() {
      chart.destroy();
    }
  };
};
