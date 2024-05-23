import React, { useState } from 'react';
import SearchBar from './components/SearchBar';
import Filters from './components/Filters';
import LogList from './components/LogList';
import DateList from './components/DateList';

const App = () => {
  const [logs, setLogs] = useState([]); // State to store logs

  const handleSearch = async (textQuery) => {
    try {
      const response = await fetch(`http://localhost:3000/queryLogs?textQuery=${textQuery}`);
      if (!response.ok) {
        throw new Error('Failed to fetch logs');
      }
      const logsData = await response.json();
      setLogs(logsData);
    } catch (error) {
      console.error('Error querying logs:', error);
    }
  };
  
  const handleFilters = async (filterValues) => {
    const queryString = new URLSearchParams(filterValues).toString();
    try {
      const response = await fetch(`http://localhost:3000/queryLogs?${queryString}`);
      if (!response.ok) {
        throw new Error('Failed to fetch logs');
      }
      const logsData = await response.json();
      setLogs(logsData);
    } catch (error) {
      console.error('Error querying logs:', error);
    }
  };

  const handleDateRange = async (filterValues) => {
    const queryString = new URLSearchParams(filterValues).toString();
    try {
      const response = await fetch(`http://localhost:3000/queryLogs?${queryString}`);
      if (!response.ok) {
        throw new Error('Failed to fetch logs');
      }
      const logsData = await response.json();
      setLogs(logsData);
    } catch (error) {
      console.error('Error querying logs:', error);
    }
  };
  

  return (
    <div>
      <h1>Log Ingestor</h1>
      <SearchBar onSearch={handleSearch} />
      <Filters onFilterChange={handleFilters} />
      <DateList onDateRange={handleDateRange}/>
      <LogList logs={logs} />
    </div>
  );
};

export default App;
