import React, { useState } from 'react'

const AddVacancyForm = () => {
	const [formData, setFormData] = useState({
		firstName: '',
		secondName: '',
		nameJob: '',
	})

	const handleSubmit = e => {
		e.preventDefault()
		fetch('http://localhost:8080/vacancyadd', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(formData),
		})
			.then(response => response.json())
			.then(data => {
				alert('Вакансия добавлена: ' + JSON.stringify(data))
				window.location.reload()
			})
			.catch(error => console.error('Ошибка при добавлении:', error))
	}

	return (
		<form onSubmit={handleSubmit} className='mb-6'>
			<h2 className='text-xl font-semibold mb-4'>Добавить вакансию</h2>
			<div className='space-y-3'>
				<input
					type='text'
					placeholder='Имя'
					value={formData.firstName}
					onChange={e =>
						setFormData({ ...formData, firstName: e.target.value })
					}
					className='w-full p-2 border rounded'
				/>
				<input
					type='text'
					placeholder='Фамилия'
					value={formData.secondName}
					onChange={e =>
						setFormData({ ...formData, secondName: e.target.value })
					}
					className='w-full p-2 border rounded'
				/>
				<input
					type='text'
					placeholder='Должность'
					value={formData.nameJob}
					onChange={e => setFormData({ ...formData, nameJob: e.target.value })}
					className='w-full p-2 border rounded'
				/>
			</div>
			<button
				type='submit'
				className='w-full bg-blue-500 text-white p-2 rounded mt-4 hover:bg-blue-600'
			>
				Добавить
			</button>
		</form>
	)
}

export default AddVacancyForm
