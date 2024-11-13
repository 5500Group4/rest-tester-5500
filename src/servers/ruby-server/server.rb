require 'sinatra'
require 'json'

set :port, 5003

before do
  response.headers['Access-Control-Allow-Origin'] = '*'
  response.headers['Access-Control-Allow-Methods'] = 'GET, POST, PUT, PATCH, DELETE, OPTIONS'
  response.headers['Access-Control-Allow-Headers'] = 'Content-Type'
end

# Handle preflight requests for CORS
options '*' do
  response.headers['Access-Control-Allow-Origin'] = '*'
  response.headers['Access-Control-Allow-Methods'] = 'GET, POST, PUT, PATCH, DELETE, OPTIONS'
  response.headers['Access-Control-Allow-Headers'] = 'Content-Type'
  200
end

users = []
next_id = 1

# Fetch all users
get '/users' do
  content_type :json
  users.to_json
end

# Fetch a user by ID
get '/users/:id' do
  content_type :json
  user = users.find { |u| u[:id] == params[:id].to_i }
  user ? user.to_json : halt(404, { error: 'User not found' }.to_json)
end

# Add a new user
post '/users' do
  content_type :json
  data = JSON.parse(request.body.read)
  if data['name'] && !data['name'].empty?
    user = { id: next_id, name: data['name'], hoursWorked: 0 }
    users << user
    next_id += 1
    status 201
    user.to_json
  else
    halt 400, { error: 'Name is required and must be non-empty' }.to_json
  end
end

# Update a user's name by ID
put '/users/:id' do
  content_type :json
  user = users.find { |u| u[:id] == params[:id].to_i }
  if user
    data = JSON.parse(request.body.read)
    if data['name'] && !data['name'].empty?
      user[:name] = data['name']
      user.to_json
    else
      halt 400, { error: 'Name is required and must be non-empty' }.to_json
    end
  else
    halt 404, { error: 'User not found' }.to_json
  end
end

# Update hours worked for a user
patch '/users/:id' do
  content_type :json
  user = users.find { |u| u[:id] == params[:id].to_i }
  if user
    data = JSON.parse(request.body.read)
    if data['hoursToAdd'].is_a?(Numeric)
      user[:hoursWorked] += data['hoursToAdd']
      user.to_json
    else
      halt 400, { error: 'Invalid hoursToAdd value' }.to_json
    end
  else
    halt 404, { error: 'User not found' }.to_json
  end
end

# Delete a user by ID
delete '/users/:id' do
  content_type :json
  user = users.find { |u| u[:id] == params[:id].to_i }
  if user
    users.delete(user)
    user.to_json
  else
    halt 404, { error: 'User not found' }.to_json
  end
end

# Delete all users
delete '/users' do
  users.clear
  next_id = 1
  [].to_json  # Return an empty array to satisfy the test expectation
end
