import React, { useEffect, useState } from 'react'

const VacancyList = () => {
	const [vacancies, setVacancies] = useState([])

	useEffect(() => {
		fetch('http://localhost:8080/vacancies')
			.then(response => response.json())
			.then(data => setVacancies(data))
			.catch(error => console.error('Ошибка при загрузке вакансий:', error))
	}, [])

	return (
		<div className='mt-6'>
			<h2 className='text-xl font-semibold mb-4'>Список вакансий</h2>
			<ul className='space-y-3'>
				{vacancies.map(vacancy => (
					<li key={vacancy.id} className='p-4 bg-white shadow rounded-lg'>
						<div className='flex justify-between items-center'>
							<div>
								<p className='font-medium'>
									{vacancy.firstName} {vacancy.secondName}
								</p>
								<p className='text-sm text-gray-600'>{vacancy.nameJob}</p>
							</div>
							<button
								onClick={() => {
									fetch(`http://localhost:8080/vacancy`, {
										method: 'DELETE',
										headers: {
											'Content-Type': 'application/json',
										},
										body: JSON.stringify({ id: vacancy.id }),
									})
										.then(() => {
											alert('Вакансия удалена')
											window.location.reload()
										})
										.catch(error =>
											console.error('Ошибка при удалении:', error)
										)
								}}
								className='bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600'
							>
								Удалить
							</button>
						</div>
					</li>
				))}
			</ul>
		</div>
	)
}

export default VacancyList
