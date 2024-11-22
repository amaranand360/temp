const axios = require('axios');

async function getCityAndCondition() {
  const url = 'https://quest.squadcast.tech/api/ENG21CS0030/weather';
  try {
    const response = await axios.get(url);
    return response.data;
  } catch (error) {
    console.error("Error fetching city and condition:", error);
    throw error;
  }
}

async function getWeatherDetails(city) {
  const url = `https://quest.squadcast.tech/api/ENG21CS0030/weather/get?q=${city}`;
  try {
    const response = await axios.get(url);
    return response.data;
  } catch (error) {
    console.error(`Error fetching weather data for ${city}:`, error);
    throw error;
  }
}

function determineBetterCity(city1, city2, weather1, weather2, condition) {
  switch (condition) {
    case 'hot':
      return weather1.temperature > weather2.temperature ? city1 : city2;
    case 'cold':
      return weather1.temperature < weather2.temperature ? city1 : city2;
    case 'windy':
      return weather1.wind > weather2.wind ? city1 : city2;
    case 'rainy':
      return weather1.rain > weather2.rain ? city1 : city2;
    case 'sunny':
      return weather1.cloud < weather2.cloud ? city1 : city2;
    case 'cloudy':
      return weather1.cloud > weather2.cloud ? city1 : city2;
    default:
      throw new Error(`Invalid condition: ${condition}`);
  }
}

async function main() {
  try {
    const cityCondition = await getCityAndCondition();
    const { City1, City2, Condition } = cityCondition;

    const weather1 = await getWeatherDetails(City1);
    const weather2 = await getWeatherDetails(City2);

    const betterCity = determineBetterCity(City1, City2, weather1, weather2, Condition);

    console.log(`The better city based on the condition '${Condition}' is: ${betterCity}`);
  } catch (error) {
    console.error("Error determining better city:", error);
  }
}

main();