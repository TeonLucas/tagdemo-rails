Rails.application.routes.draw do
  get 'appointments/index'
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  resources :dataentry, only: [:index]
  resources :images, only: [:index]
  resources :goserver, only: [:index]
end
