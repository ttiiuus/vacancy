import React from 'react'
import AddVacancyForm from './components/AddVacancyForm'
import VacancyList from './components/VacancyList'

const App = () => {
	return (
		<div className='container mx-auto p-4 max-w-2xl'>
			<h1 className='text-2xl font-bold text-center mb-6'>
				Управление вакансиями
			</h1>
			<AddVacancyForm />
			<VacancyList />
		</div>
	)
}

export default App
