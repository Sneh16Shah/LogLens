import React, { useState } from 'react';

const DateList = ({ onDateRange }) => {
  const [dateRange, setDateRange] = useState({
    startDate: '',
    endDate: '',
  });

  const handleDateRanges=(e)=>{
    const {name, value} = e.target
    setDateRange({...dateRange, [name]:value})
  }
  const handleDateSearch = () => {
    onDateRange(dateRange);
  };

  return (
    <div>
      <input
        type="text"
        placeholder="Start Date(YYYY-MM-DD)"
        name="startDate"
        value={dateRange.startDate}
        onChange={handleDateRanges}
      />
      <input
        type="text"
        placeholder="End Date(YYYY-MM-DD)"
        name="endDate"
        value={dateRange.endDate}
        onChange={handleDateRanges}
      />
      <button onClick={handleDateSearch}>Date Range Search</button>
    </div>
  );
};

export default DateList;
