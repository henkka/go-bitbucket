package bitbucket

import (
	"context"
)

type GroupsService service

const groupsApiName = "api"

type GroupMember struct {
	Name                        string            `json:"name"`
	EmailAddress                string            `json:"emailAddress"`
	Active                      bool              `json:"active"`
	DisplayName                 string            `json:"displayName"`
	Type                        string            `json:"type"`
	DirectoryName               string            `json:"directoryName"`
	Deletable                   bool              `json:"deletable"`
	LastAuthenticationTimestamp int64             `json:"lastAuthenticationTimestamp"`
	MutableDetails              bool              `json:"mutableDetails"`
	MutableGroups               bool              `json:"mutableGroups"`
	Links                       map[string][]Link `json:"links,omitempty"`
}

type GroupMemberList struct {
	ListResponse
	GroupMembers []*GroupMember `json:"values"`
}

type GroupMemberSearchOptions struct {
	ListOptions
	GroupName string `url:"context"`
}

func (s *GroupsService) ListGroupMembers(ctx context.Context, opts *GroupMemberSearchOptions) ([]*GroupMember, *Response, error) {
	g := "admin/groups/more-members"
	var l GroupMemberList
	resp, err := s.client.GetPaged(ctx, groupsApiName, g, &l, opts)
	if err != nil {
		return nil, resp, err
	}
	return l.GroupMembers, resp, nil
}
